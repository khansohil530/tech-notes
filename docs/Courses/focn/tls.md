# TLS

Earlier application protocols (like HTTP, FTP, SMTP) used to send data in plain text which could be intercepted, read
and altered by any middle man on network. Bad agent could use this to impersonate as server
and get your secret data like password, which is unacceptable specially for online banking and e-commerce websites. 
The earliest attempt to solve this problem was by **SSL** (Secure Socket Layer) protocol which laid the foundation for 
the standard security protocol used now, called **TLS** (Transport Layer Security). It can be added as an extension
over existing protocol to provide a secure channel for communication, by providing ways to authenticate server while
maintaining integrity and confidentiality of data in transit (essential for developing trusted web APIs).

## TLS1.2

TLS 1.2 was the first major version of TLS which made web communication secure by providing,

1. Encryption, to make data is unreadable to others.
2. Authentication, to verify the identity of communicating server.
3. Integrity, so that data can’t be altered silently.

For example, HTTP secures its communication using TLS, making it HTTPS. However, the browser still sends HTTP requests, 
but before sending this HTTP request to lower/upper layer, TLS encrypts/decrypt the information using established 
cryptographic parameters. These parameters are decided by both parties after TCP connection is established, over a bunch of request/response
together known as **TLS handshake**. The handshake usually involves following steps, which will discuss further to explain
the working of TLS to make communication secure.

```mermaid
--8<-- "docs/Courses/focn/diagram/tls_rsa_handshake.mmd"
```

HTTPS uses **Asymmetric Encryption**(1) for generating the session keys, which is used to encrypt the rest of
communication as **Symmetric Encryption**(2), but the specific algorithm is decided from
the supported Cipher Suite shared during Client and Server Hello request. 
The `ClientHello` usually lists the supported TLS specs (3), while the `ServerHello` decides the specific implementation.
This allows clients to list their possible options, while server can choose the level of security as per its requirement.   
{.annotate}

1. uses two different keys each for encryption (_private key_ `SK`) and decryption (_private key_ `PK`). Operations are
   more expensive than Symmetric encryption algorithms, as such its only used for encrypting smaller dataset, like 
   metadata or headers. Algorithms developed on this idea: RSA, ECDHE.
2. encrypt and decrypt using same secret key. Uses less expensive computation, as such it's much faster for encrypting 
   larger dataset. Algorithms like AES, ChaCha20 implements this category of encryption.
3. like the Key Exchange Algorithm ([RSA](#rsa)/[DHE](#diffie-hellman)), Authentication Algorithm (RSA/ECDSA), 
   Symmetric Encryption Algorithm (AES/ChaCha20), Encryption Mode (GCM/POLY), Integrity Algorithm (SHA)


### RSA

In above TLS handshake, we've used RSA key exchange algorithm to generate pre-master secret which is used to generate 
session keys for encryption. For a better understanding, check following steps:

```mermaid
--8<-- "docs/Courses/focn/diagram/rsa_key_exchange.mmd"
```

The above diagram is a little simplified to avoid cluttering with implementation details. For example, to generate
secure session keys (1), each host runs a **PRF** (2) over Pre-Master Secret, Client Random and Server Random parameters.
Using **Client Random** and **Server Random** ensure that every TLS session produces unique, unpredictable 
keys even if the same secrets or certificates are reused. This provides security against replay and precomputation 
attacks.
{.annotate}

1. like write key, MAC keys, IVs. 
2. Pseudo-Random Function, used to generate random but deterministic output.


??? note "Limitation: No Forward Secrecy"
    RSA doesn't provide any **forward secrecy**. If any parties RSA private key is ever compromised, attacker
    could decrypt the pre-master secret shared over network, and generate same symmetric key. This is a major security 
    failure, as they could now decrypt all the past recorded exchange, which is why RSA key exchange isn't used by any
    secure protocol anymore. It was resolved using [Diffie-Hellman](#diffie-hellman) algorithm.

---

### Certificates

The above design is still prone to man-in-the-middle attack unless we can verify that the Public Key received is same
as the one sent by Server. For example,

```mermaid
--8<-- "docs/Courses/focn/diagram/tls_cert_mitm_attack.mmd"
```

Certificates help prevent this by allowing servers to embed their public key in a document signed by a trusted CA 
(Certificate Authority), which can be verified by Client for its authenticity and integrity. This way, if attacker 
makes any change to public key in certificate, the client can easily detect it and abort the connection. The general 
gist is as follows:

```mermaid
--8<-- "docs/Courses/focn/diagram/cert_verification.mmd"
```

To make the Public Key verifiable with an owner, the CA creates a digital document storing server's domain and public 
key, along with other required information and attaching its signature. This signature is created using the content of 
certificate, which are hashed and encrypted using CA's private key. Below is a simplified view of `X.509` certificate
with essential fields required for verification of certificate.
```text
┌──────────────────────────────────────────────────┐
│                X.509 CERTIFICATE                 │
├──────────────────────────────────────────────────┤
│ ...                                              │
│ Issuer:                                          │
│   CN = Trusted Intermediate CA                   │
│   O  = Example CA Inc.                           │
│                                                  │
│ Validity:                                        │
│   Not Before:  2025-01-01 00:00:00 UTC           │
│   Not After :  2026-01-01 23:59:59 UTC           │
│                                                  │
│ Subject:                                         │
│   CN = example.com                               │
│                                                  │
│ Subject Public Key Info:                         │
│   Algorithm: RSA                                 │
│   Public Key: ...                                │
│   ...                                            │
├──────────────────────────────────────────────────┤
│ Certificate Signature:                           │
│   Algorithm: sha256WithRSAEncryption             │
│   Signature: ...                                 │
└──────────────────────────────────────────────────┘
```

When the client browser receives the certificate, it verifies certificates authenticity by building a
**Certificate Chain**(1) using **Trust Store**(2) of its systems. For chaining, the server sends one or more 
intermediate CA certificates which are verifiable by root CA along its own certificate. Using this, 
the browser can link the Issuer and Subject fields as follows: `server cert -issued by -> Intermediate CA -issued by -> 
Root CA` to create a certificate chain. For each link in the chain, the certificate is verified by matching the signature decrypted using 
issuer's public key against the recomputed hash of certificate data. This process is done on every step of chaining 
until Root CA, which is considered trusted own its own (self-signed). The whole verification can fail at any point in
chain, if the browser can't find the issuer or verify the signature. Once verified, browser does additional checks like 
domain name, validity dates, key usage, revocation status, etc. before continuing TLS handshake.
{.annotate}

1. sequence of certificates where each certificate is signed by the next one above it, ending at a root CA that the 
   browser already trusts.
2. every browser and operating system ships with a built-in list of trusted Root CA certificates.

Following this procedure both hosts can agree on their public parameters with integrity and authenticity, which are then
further used for generating session keys. Either host send a `ChangeCipherSec` message, it indicates the start of
encrypted communication using session keys. 

??? note "Use Cryptographic Libraries"
    Implementing the algorithms used for encryption/decryption from scratch is generally avoided because it's extremely 
    easy to make subtle mistakes that break security, even if the math is correct. To ensure the security, these
    algorithms have requirements like correct composition of primitives, secure randomness, constant-time operations,
    etc. To handle these details safely, well-established libraries (like `OpenSSL`) have been developed and tested by
    subject experts. It's recommended to use such existing libraries which significantly reduces the risk of
    vulnerabilities, errors and allows developers to rely on proven, standardized implementations.


## TLS1.3

One of the major limitations of TLS1.2 was it didn't support forward secrecy, due to default usage of RSA key exchange
algorithm. Additionally, it took 2 Round trips to complete the handshake which incurred extra latency for new 
connections, used weak cryptographic algorithms and parameters (like SHA-1, CBC-mode) which could silently downgrade 
security if misconfigured. 
??? note "Example: Forward Secrecy"
    ```mermaid
    --8<-- "docs/Courses/focn/diagram/forward_secrecy.mmd"
    ```

TLS1.3 solved all of these issues by firstly removing weaker cryptographic options and placing secure default inplace 
(1) to make it impossible to deploy weaker crypto accidentally. Secondly, it made **Forward Secrecy** mandatory by using
ephemeral key exchange for every connection. And finally it made the handshake faster by reducing it to 1 Round trip 
which mattered a lot for high-latency devices. The complete handshake is as follows: 
{.annotate}

1. Like AEAD ciphers, ECDHE, SHA-256/384 

```mermaid
--8<-- "docs/Courses/focn/diagram/tls_1_3_handshake.mmd"
```

!!! note "RTT"
    One RTT is `Client sends -> Server replies -> Client receives` or to simplify it's the amount of time the client
    must wait before it can start sending application data. For above case, starts in the same flight as its Finished 
    message. 

Firstly, let's understand Diffie-Hellman key exchange algorithm which enables this version of TLS.

### Diffie Hellman

The core of this key exchange algorithm is following mathematics formulas: 

- Given common group parameters (`g`, `p`) where `g` is a generator (1) value, `p` is a large prime number and two
  private values for both parties (`a`, `b`). 
  {.annotate}

      1. number which produces large, unpredictable set of exponential results, to avoid cycling through only a small subset 
         of values which would leak information and weaken security.

- You can send over following Key Share values over network: $A = g^a \mod p$ and $B = g^b \mod p$
- And generate shared secret on both side: $S = B^a = A^b$

This is secure because computing $a$ or $b$ from given $A, B, g, p$ requires solving the discrete logarithm problem(1) which
is computationally infeasible (with proper parameters) at the time . Finally, when both parties shared secret, it can
be passed into KDF (Key Derivation Function) to generate secure session keys to continue communication over an encrypted
channel. 
{.annotate}

1. math problem that's easy to compute in one direction, but extremely hard to reverse.

This algorithm is the classic `mod-p DH`, but its math was computationally expensive. This was later optimized to use 
elliptic curves which is faster to computer, and was named as **Elliptic Curve DH** (`ECDH`). Also, it's important to 
discard `a` and `b` right after the handshake, and since they're ephemeral the algorithm is named to **ECDH Ephemeral**
(`ECDHE`). This way, every session uses generates different session keys which are discarded as soon as connection is 
closed, allowing it to maintain **forward secrecy**.

---
Until now, we've covered `ECDHE` which generates shared secret between both client and server. 
Now we need to authenticate the server, which is using the Certificate (as discussed [before](#certificates)) but this
time there's a key difference. 

The Certificate only ensures that the public key mentioned belongs to given domain, it doesn't ensure that the server 
is in possession of respective private key. For TLS1.2, this possession check was done implicitly either by decrypting the
premaster secret when TLS is configured with RSA or when verifying the contents of `ServerKeyExchange` message signed 
by server's private key when configured for `ECDHE`. Without confirming the possession of private key, attacker can 
replay the valid certificate (since cert is public) to impersonate as its owner during the handshake. To verify possession
of private key, TLS1.3 uses **CertificateVerify** message explicitly, which sends a signature generates using its 
authentication private key and a hash of its current handshake. This signature can be verified by client by decrypting
it with the public key in the certificate and matching it against the hash of current handshake.

!!! note ""
    Certificates are long-lived while DH keys are ephemeral. So the private key mentioned in CertificateVerify is the 
    one used to authenticate the server and it long-lived, and not the one which will encrypt the session data.    

---

The **Encrypted Extensions** message is used to send other protocol extensions that affect the connection and aren't 
part of TLS  handshake. For example, ALPN (Application Layer Protocol Negotiation) to decide b/w HTTP/2 or HTTP/1.1. 
Since these values influence how the application protocol behaves, they're encrypted to prevent fingerprinting(1) and 
downgrade attacks (2).
{.annotate}

1. attacker analyzes network traffic patterns of encrypted connections to identify the specific website a user is
   visiting, essentially creating a unique "fingerprint" for that site's traffic to de-anonymize users.
2. attack that forces to use weaker, outdated security protocols or versions

---

Finally, both parties ensure that they've derived the same handshake secrets from the same handshake transcript, 
without interference. Without verifying this, neither side knows about the incorrect handshake until the data breaks. 

To achieve this, both client and server send a **Finished** message, which is computed as an HMAC(1) over the hash of the
entire handshake transcript, using a finished key (client_finished_key or server_finished_key) derived from the DH
shared secret.
{.annotate}

1. HMAC (Hash-based Message Authentication Code) verifies both the integrity and authenticity of data by combining a 
   cryptographic hash function (e.g., SHA-256) with a secret key. 

The peer recomputes the handshake transcript hash and verifies the MAC, which provides explicit key confirmation and integrity
protection for the entire handshake. This proves that both parties participated correctly in the key exchange and that
no attacker modified, reordered, or removed any handshake message.
