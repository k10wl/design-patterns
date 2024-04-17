class Facade {
  #subsystem: LowLevelLogic;
  constructor(subsystem?: LowLevelLogic) {
    this.#subsystem = subsystem ?? new LowLevelLogic();
  }
  pickDressing(): string {
    if (this.#subsystem.measureTemperature() > 0) {
      return "jacket";
    } else {
      return "coat";
    }
  }
}

class LowLevelLogic {
  measured: boolean = false;
  public measureTemperature(): number {
    this.measured = !this.measured;
    return this.measured ? 1 : -1;
  }
}

function clientCode(facade: Facade) {
  console.log(facade.pickDressing());
}

const facade = new Facade();
clientCode(facade);
clientCode(facade);
clientCode(facade);
clientCode(facade);
