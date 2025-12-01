# Database Security

Data is a critical part of every business which should be handled responsibly for successfully running
a business. Security plays a vital role in managing such data by using set of policies and controls to
protect data from unauthorized access, misuse, alteration or destruction ensuring only the right user
can access the right data. Key goals of security in DB involves:

- Keeping data confidential, by protecting it from unauthorized access
- Maintaining data integrity, by controlling access so that it can't be altered by just anybody.
- Ensuring data is available whenever needed, it shouldn't be lost on accidents or crashes.

Few components which helps DB achieve these goals

- Access Control: Control who can access what data and how they can use it by using proper authentication and
  authorization. For example, only admins should've access to creating and deleting tables.
- Encryption: Data should be protected from 3rd parties by using safe encryption at both rest and transit to 
  keep them confidential.
- Backup & Recovery: Keeping regular backups and redundant copies can prevent data loss due to failure, or crashes.

REST apps sometimes requires database tables to be present, and it’s usually covered with the startup 
of app to create the table if not present. This is a bad practice and should be avoided,
because when your app users are interacting with your database — they’ll have full privilege to your 
DB which can cause serious harm like SQL Injections or XSS attack to drop the table. 
Instead, you should keep separate users for creating tables, schema, etc. and have separate users for
read/write permissions. You can also maintain separate connection pools in client side for each of these 
read/update/delete operations.

## Homomorphic Encryption

Encryption is transforming data into random text which doesn’t make sense when looked at. 
To make the sense out of it, you’ll have to decrypt the random text. 
You can do this in two different way, by using a common key (symmetric encryption) for both encryption
and decryption or by using separate keys (asymmetric encryption) for encryption and decryption.

In database systems, we can’t simply encrypt the data because of few reasons
- queries need plain data to perform their work. To work with encrypted data, it’ll have to decrypt it first 
  everytime which isn’t optimal.
- Analysis of data, indexing, tuning the data needs plain text
- application needs recognizable data to process it.
- Layer 7 reverse proxies terminates TLS connections so that it can read the traffic and apply
  routing rules on it.

*Homomorphic Encryption* allows you to do all these operations on encrypted data by allowing 
Arithmetic operations on encrypted data. You can use indexes and query on encrypted data using these
Arithmetics because at low level each of these operations are simply performed by comparison, shifting bits,
adding, etc. However, it’s still in PoC as actual querying is too slow for production system.