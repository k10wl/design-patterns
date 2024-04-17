interface Component {
  notify(): string;
}

class BaseNotification implements Component {
  notify(): string {
    return "base notification";
  }
}

class Decorator implements Component {
  component: Component;
  constructor(component: Component) {
    this.component = component;
  }
  public notify(): string {
    console.log("> calling " + this.constructor.name);
    return this.component.notify();
  }
}

class TelegramDecorator extends Decorator {
  public notify(): string {
    return `telegram(${super.notify()})`;
  }
}

class SlackDecorator extends Decorator {
  public notify(): string {
    return `slack(${super.notify()})`;
  }
}

const base = new BaseNotification();
const tg = new TelegramDecorator(base);
const slack = new SlackDecorator(tg);

console.log("[result]:", slack.notify());
