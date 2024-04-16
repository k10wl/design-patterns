class Singleton {
  static instance: Singleton;
  constructor() {}
  public static getInstance(): Singleton {
    if (!this.instance) {
      console.log("> no instance, creating new");
      this.instance = new Singleton();
    } else {
      console.log("> instance exists, returning existing");
    }
    return this.instance;
  }
}

for (let i = 0; i < 10; i++) {
  Singleton.getInstance();
}
