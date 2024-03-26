class Target {
  public request(): string {
    return "Default target behaviour";
  }
}

class LegacyClass {
  public specificRequest(): string {
    return ".eetpadA eht fo roivaheb laicepS";
  }
}

class LegacyAdapter extends Target {
  private adaptee: LegacyClass;

  constructor(adaptee: LegacyClass) {
    super();
    this.adaptee = adaptee;
  }

  public request(): string {
    const res = this.adaptee.specificRequest();
    return res.split("").reverse().join("");
  }
}

function clientCode(target: Target): void {
  console.log(target.request());
}

console.log("> Target with default output");
const target = new Target();
console.log(target.request());

console.log();

console.log("> Legacy called directly");
const legacy = new LegacyClass();
console.log(legacy.specificRequest());

console.log();

console.log(
  "> Concrete legacy decoding moved to adapter implementation, client does not rely on specific code, only on interface"
);
const adapter = new LegacyAdapter(legacy);
clientCode(adapter);
