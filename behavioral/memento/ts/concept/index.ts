class Originator {
  name: string;
  constructor(name: string) {
    this.name = name;
  }
  saveState(): ConcreteMemento {
    return new ConcreteMemento(this.name);
  }
  restoreState(memento: Memento): void {
    this.name = memento.getData();
  }
}

interface Memento {
  getData(): string;
}

class ConcreteMemento implements Memento {
  #data: string;
  constructor(data: string) {
    this.#data = data;
  }
  getData(): string {
    return this.#data;
  }
}

class Caretaker {
  #history: Memento[];
  originator: Originator;
  constructor(originator: Originator) {
    this.#history = [];
    this.originator = originator;
  }
  update(name: string): void {
    this.#history.push(this.originator.saveState());
    this.originator.name = name;
  }
  prev(): void {
    this.originator.restoreState(this.#history.pop()!);
  }
}

const originator = new Originator("-1");
const careTaker = new Caretaker(originator);
const iterations = 5;
console.log(originator.name);
for (let i = 0; i < iterations; i++) {
  careTaker.update(i.toString(10));
  console.log(originator.name);
}
for (let i = 0; i < iterations; i++) {
  careTaker.prev();
  console.log(originator.name);
}
