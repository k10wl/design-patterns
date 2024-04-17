interface Composite {
  print(offset?: number): void;
}

class Leaf implements Composite {
  text: string;
  constructor(text: string) {
    this.text = text;
  }
  print(offset: number = 0): void {
    console.log(this.text.padStart(offset + this.text.length));
  }
}

class Branch implements Composite {
  name: string;
  leafs: Leaf[];
  branches: Branch[];
  constructor(name: string) {
    this.name = name;
    this.leafs = []
    this.branches = []
  }
  addBranch(name: string): Branch {
    const branch = new Branch(name)
    this.branches.push(branch);
    return branch
  }
  addLeaf(name: string): Leaf {
    const leaf = new Leaf(name)
    this.leafs.push(leaf);
    return leaf
  }
  print(offset: number = 0): void {
    console.log(this.name.padStart(offset + this.name.length));
    for (let i = 0; i < this.branches.length; i++) {
      this.branches[i].print(offset + 2);
    }
    for (let i = 0; i < this.leafs.length; i++) {
      this.leafs[i].print(offset + 2);
    }
  }
}

const root = new Branch("root");
for (let i = 0; i < 3; i++) {
  const branch = root.addBranch(`branch #${i}`);
  for (let j = 0; j < 3; j++) {
    branch.addLeaf(`leaf #${i}`)
  }
}

function clientCode(structure: Composite) {
  structure.print()
}

clientCode(root)
