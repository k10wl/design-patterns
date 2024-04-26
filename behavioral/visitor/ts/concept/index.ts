interface Component {
  accept(visitor: IVisitor): void;
}

class Stove implements Component {
  name = "Retro Burner";
  accept(visitor: IVisitor) {
    visitor.visitStove(this);
  }
}

class TV implements Component {
  brand = "MegaByte";
  accept(visitor: IVisitor) {
    visitor.visitTV(this);
  }
}

class Radio implements Component {
  freequency = "96.24fm";
  accept(visitor: IVisitor) {
    visitor.visitRadio(this);
  }
}

interface IVisitor {
  visitStove: (stove: Stove) => void;
  visitTV: (stove: TV) => void;
  visitRadio: (stove: Radio) => void;
}

class NameVisitor implements IVisitor {
  visitStove(stove: Stove) {
    console.log("Commmon stove name:", stove.name);
  }
  visitTV(stove: TV) {
    console.log("Commmon TV name:", stove.brand);
  }
  visitRadio(stove: Radio) {
    console.log("Commmon radio name: ", stove.freequency);
  }
}

class MatchesNameVisitor implements IVisitor {
  #regex = /(a|o|i|y|u|e)/;
  visitStove(stove: Stove) {
    console.log("Stove matches:", this.#regex.test(stove.name));
  }
  visitTV(stove: TV) {
    console.log("TV matches:", this.#regex.test(stove.brand));
  }
  visitRadio(stove: Radio) {
    console.log("Radio matches: ", this.#regex.test(stove.freequency));
  }
}

const stove = new Stove();
const tv = new TV();
const radio = new Radio();
const components = [stove, tv, radio];

const nameVisitor = new NameVisitor();
const matchesName = new MatchesNameVisitor();

function clientCode(components: Component[], visitor: IVisitor): void {
  for (let i = 0; i < components.length; i++) {
    components[i].accept(visitor);
  }
}

clientCode(components, nameVisitor);
clientCode(components, matchesName);
