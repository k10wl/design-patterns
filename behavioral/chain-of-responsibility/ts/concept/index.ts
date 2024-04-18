class CustomRequest {
  command: string;
  isAdmin: boolean;
  isPremium: boolean;
  constructor(command: string, isAdmin: boolean, isPremium: boolean) {
    this.command = command;
    this.isAdmin = isAdmin;
    this.isPremium = isPremium;
  }
}

abstract class Handler {
  private next: Handler | undefined = undefined;
  constructor() {
    this.next = undefined;
  }
  public setNext(handler: Handler): Handler {
    this.next = handler;
    return handler;
  }
  public handle(request: CustomRequest): void {
    this.next?.handle(request);
  }
}

class PremiumHandler extends Handler {
  public handle(request: CustomRequest): void {
    if (!request.isAdmin) {
      console.log("> not premium, interrupting the chain");
      return;
    }
    console.log("> is premium");
    super.handle(request);
  }
}

class AdminHandler extends Handler {
  public handle(request: CustomRequest): void {
    if (!request.isAdmin) {
      console.log("> not admin, interrupting the chain");
      return;
    }
    console.log("> is admin");
    super.handle(request);
  }
}

class CounterHandler extends Handler {
  public handle(request: CustomRequest): void {
    console.log("> checking if there are enough cookies...");
    super.handle(request);
  }
}

class CommandHandler extends Handler {
  public handle(request: CustomRequest): void {
    console.log("> executing:", request.command);
    super.handle(request);
  }
}

function clientCode(handler: Handler) {
  console.log("[req]");
  handler.handle(new CustomRequest("bring_cookies", true, true));
  console.log("[req]");
  handler.handle(new CustomRequest("bring_cookies", false, true));
}

const premiumHandler = new PremiumHandler();
const adminHandler = new AdminHandler();
const counterHander = new CounterHandler();
const commandHandler = new CommandHandler();

premiumHandler
  .setNext(adminHandler)
  .setNext(counterHander)
  .setNext(commandHandler);

clientCode(premiumHandler);
clientCode(counterHander);
