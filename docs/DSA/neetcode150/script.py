import re
from pathlib import Path

# -------- CONFIG --------
MARKDOWN_FILE = "index.md"  # input markdown file
PY_DIR = Path("src/py")
GO_DIR = Path("src/go")
# ------------------------

PROBLEM_CELL_RE = re.compile(r"<td>(.*?)</td>", re.DOTALL)
STAR_RE = re.compile(r":star:\{\.star\}")

def normalize_filename(name: str) -> str:
    name = STAR_RE.sub("", name)         # remove star markers
    name = name.lower()
    name = name.replace("&", "and")
    name = name.replace("/", "_")
    name = re.sub(r"[^a-z0-9\s_]", "", name)
    name = re.sub(r"\s+", "_", name.strip())
    return name

def extract_problem_names(markdown: str):
    cells = PROBLEM_CELL_RE.findall(markdown)
    problems = []

    for cell in cells:
        # skip non-problem columns
        if "badge" in cell or "notes" in cell or "site" in cell:
            continue

        # ignore difficulty / reference columns
        text = re.sub(r"<.*?>", "", cell).strip()
        if not text or len(text.split()) < 2:
            continue

        problems.append(text)

    return problems

def touch(path: Path, content: str = ""):
    if not path.exists():
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(content, encoding="utf-8")

def main():
    md_path = Path(MARKDOWN_FILE)
    if not md_path.exists():
        raise FileNotFoundError(f"{MARKDOWN_FILE} not found")

    markdown = md_path.read_text(encoding="utf-8")
    problems = extract_problem_names(markdown)

    created = 0

    for problem in problems:
        filename = normalize_filename(problem)

        md_file = Path(f"{filename}.md")
        py_file = PY_DIR / f"{filename}.py"
        go_file = GO_DIR / f"{filename}.go"

        touch(md_file, f"# {problem}\n")
        touch(py_file, f"# {problem}\n")
        touch(go_file, f"// {problem}\n")

        created += 1

    print(f"✔ Processed {created} problems")

if __name__ == "__main__":
    main()
