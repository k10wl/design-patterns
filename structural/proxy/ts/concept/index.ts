class Original {
  public foo(): void {
    console.log("public foo")
  }
  public bar(): void {
    console.log("public bar")
  }
  public baz(): void {
    console.log("public baz")
  }
}

interface IProxy {
  foo(): void
}

class OriginProxy implements IProxy {
  #origin: Original
  constructor(origin: Original) {
    this.#origin = origin
  }
  public foo(): void {
    console.log('> calling through proxy')
    this.#origin.foo()
  }
}

function clientCode(proxy: IProxy) {
  proxy.foo()
}

const original = new Original()
const proxy = new OriginProxy(original)
clientCode(proxy)
