class Flyweight {
  #color: string;
  constructor(color: string) {
    this.#color = color;
  }
  public operation(uniqueState: UniqueState): void {
    console.log(
      `shared: ${JSON.stringify(this.#color)};`,
      `unique: ${JSON.stringify(uniqueState)};`,
    );
  }
}

class FlyweightFactory {
  #map: Map<string, Flyweight>;
  constructor() {
    this.#map = new Map();
  }

  public getFlyweight(key: string): Flyweight {
    let cached = this.#map.get(key);
    if (!cached) {
      const created = new Flyweight(key);
      this.#map.set(key, created);
      cached = created;
      console.log("> creating", key);
    } else {
      console.log("> reusing", key);
    }
    return cached;
  }
}

class UniqueState {
  value: number;
  constructor() {
    this.value = Math.random();
  }
}

const factory = new FlyweightFactory();
for (let i = 0; i < 5; i++) {
  const unique = new UniqueState();
  const green = factory.getFlyweight("green");
  green.operation(unique);
  const green2 = factory.getFlyweight("green");
  green2.operation(unique);
}
