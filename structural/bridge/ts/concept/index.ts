interface Implementation {
  searchForDevice(): void;
  sendHandshake(): void;
  establishConnection(): boolean;
}

class UserInterface {
  #implementation: Implementation;
  constructor(implementation: Implementation) {
    this.#implementation = implementation;
  }
  connectEarphones(): void {
    console.log("[user action]: connect earphones");
    this.#implementation.searchForDevice();
    this.#implementation.sendHandshake();
    if (this.#implementation.establishConnection()) {
      console.log("[user action]: device connected!");
    }
  }
}

class SoundCore implements Implementation {
  searchForDevice(): void {
    console.log("[implementation]: searching for device");
  }
  sendHandshake(): void {
    console.log("[implementation]: sending handshake");
  }
  establishConnection(): boolean {
    console.log("[implementation]: establishing connection");
    return true;
  }
}

function userAction(lowLevel: Implementation): void {
  const ui = new UserInterface(lowLevel);
  ui.connectEarphones();
}

const lowLevel = new SoundCore();
userAction(lowLevel);
