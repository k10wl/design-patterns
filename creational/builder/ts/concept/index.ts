type TSquare = {
  sideSize: number;
  color: string;
};

type TCircle = {
  radius: number;
  color: string;
};

interface IBuidler {
  reset(): void
  setSize(): void;
  paint(): void;
}

class Circle implements TCircle {
  radius: number;
  color: string;
  constructor() {
    this.radius = 0;
    this.color = "transparent";
  }
}

class Square implements TSquare {
  sideSize: number;
  color: string;
  constructor() {
    this.sideSize = 0;
    this.color = "transparent";
  }
}

class CircleBuilder implements IBuidler {
  #circle: TCircle;
  constructor() {
    this.#circle = new Circle();
  }
  public reset(): void {
    this.#circle = new Circle();
  }
  public setSize(): void {
    this.#circle.radius = 2;
  }
  public paint(): void {
    this.#circle.color = "black";
  }
  getShape(): TCircle {
    return this.#circle;
  }
}

class SquareBuilder implements IBuidler {
  #square: TSquare;
  constructor() {
    this.#square = new Square();
  }
  public reset() {
    this.#square = new Square();
  }
  public setSize(): void {
    this.#square.sideSize = 5;
  }
  public paint(): void {
    this.#square.color = "white";
  }
  getShape(): TSquare {
    return this.#square;
  }
}

class Director {
  #buider: IBuidler;
  constructor(builder: IBuidler) {
    this.#buider = builder;
    builder.reset()
  }
  public build(): void {
    this.#buider.setSize();
    this.#buider.paint();
  }
  public setBuilder(builder: IBuidler): void {
    this.#buider = builder;
    this.#buider.reset()
  }
}


const circleBuilder = new CircleBuilder();
const director = new Director(circleBuilder)
director.build()
console.log("> circle", circleBuilder.getShape());
const squareBuilder = new SquareBuilder();
director.setBuilder(squareBuilder)
director.build()
console.log("> square", squareBuilder.getShape());
