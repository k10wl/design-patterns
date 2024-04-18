interface Mediator {
  notify(receiver: object, data: number): void;
}

class BaseComopnent {
  protected mediator: Mediator;
  constructor(mediator?: Mediator) {
    this.mediator = mediator!;
  }
  setMediator(mediator: Mediator) {
    this.mediator = mediator;
  }
}

class ComponentA extends BaseComopnent {
  logNumber(number: number): void {
    this.mediator.notify(this, number);
  }
}

class ComponentB extends BaseComopnent {
  countOrder(number: number): void {
    this.mediator.notify(this, number);
  }
}

class ConcreateMediator implements Mediator {
  #a: ComponentA;
  #b: ComponentB;
  constructor(a: ComponentA, b: ComponentB) {
    this.#a = a;
    this.#b = b;
  }
  public notify(sender: object, data: number): void {
    switch (sender) {
      case this.#a:
        console.log("react to a:", data);
        break;
      case this.#b:
        console.log("react to b:", data);
        break;
      default:
        throw new Error("unhandled mediator");
    }
  }
}

const a = new ComponentA();
const b = new ComponentB();
const mediator = new ConcreateMediator(a, b);
a.setMediator(mediator);
b.setMediator(mediator);

a.logNumber(69);
b.countOrder(420);
