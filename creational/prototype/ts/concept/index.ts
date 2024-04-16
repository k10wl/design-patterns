interface Prototype {
  clone(): this;
}

class Shape implements Prototype {
  x: number;
  y: number;
  constructor() {
    this.x = 3;
    this.y = 3;
  }
  clone(): this {
    return JSON.parse(JSON.stringify(this));
  }
}
class Circle extends Shape {
  radius: number;
  constructor() {
    super();
    this.radius = 3;
  }
}

const shapeOrigin = new Shape();
const shapeClone = shapeOrigin.clone();
console.log(
  "> shapeOrigin === shapeClone",
  shapeOrigin.x === shapeClone.x && shapeOrigin.y === shapeClone.y,
  JSON.stringify(shapeOrigin), 
  JSON.stringify(shapeClone), 
);

const circleOrigin = new Circle();
const circleClone = circleOrigin.clone();
console.log(
  "> circleOrigin === circleClone",
  circleOrigin.x === circleClone.x &&
    circleOrigin.y === circleClone.y &&
    circleOrigin.radius === circleClone.radius,
  JSON.stringify(circleOrigin), 
  JSON.stringify(circleClone), 
);
