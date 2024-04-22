interface Observer {
  notify(version: string): void;
}

interface Publisher {
  subscribe(sub: Observer): void;
  unsubscribe(sub: Observer): void;
  notify(): void;
}

class ConcretePublisher implements Publisher {
  #subscribers: Observer[];
  #version: string;

  constructor() {
    this.#subscribers = [];
    this.#version = "0.1.0";
  }

  subscribe(...observers: Observer[]): void {
    for (let i = 0; i < observers.length; i++) {
      if (this.#subscribers.includes(observers[i])) {
        console.log(
          `x [Publisher]: already subscribed ${JSON.stringify(observers[i])}`,
        );
        continue;
      }
      this.#subscribers.push(observers[i]);
      console.log(
        `+ [Publisher]: subscribed "${JSON.stringify(observers[i])}"`,
      );
    }
  }

  unsubscribe(observer: Observer): void {
    const i = this.#subscribers.indexOf(observer);
    if (i === -1) {
      console.log(`x [Publisher]: not subscribed ${JSON.stringify(observer)}`);
      return;
    }
    this.#subscribers.splice(i, 1);
    console.log(`- [Publisher]: removed "${JSON.stringify(observer)}"`);
  }

  updateVersion(version: string): void {
    this.#version = version;
    this.notify();
  }

  notify() {
    for (let i = 0; i < this.#subscribers.length; i++) {
      this.#subscribers[i].notify(this.#version);
    }
  }
}

class ConcreteObserver implements Observer {
  #knownVersion: string;
  name: string;

  constructor(name: string) {
    this.name = name;
    this.#knownVersion = "none";
  }

  notify(version: string): void {
    this.#knownVersion = version;
    console.log(
      `> [${this.name}]: updated known version to "${this.#knownVersion}"`,
    );
  }
}

const bob = new ConcreteObserver("Bob");
const annie = new ConcreteObserver("Annie");
const jack = new ConcreteObserver("Jack");

const versionControll = new ConcretePublisher();

versionControll.subscribe(bob, annie);
versionControll.updateVersion("0.2.0");
versionControll.subscribe(jack);
versionControll.updateVersion("0.2.1");
versionControll.unsubscribe(bob);
versionControll.updateVersion("0.2.2");
versionControll.unsubscribe(bob);
versionControll.subscribe(jack);
versionControll.updateVersion("0.3.0");
