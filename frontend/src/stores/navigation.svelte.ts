class NavigationStore {
  history = $state<string[]>([]);
  curr = $state(0);

  get canGoBack() {
    console.log(`cannot go back as current index is ${this.curr}`);
    return this.curr > 0;
  }

  get canGoForward() {
    console.log(`cannot go forward as current index is ${this.curr} and history length is ${this.history.length}`);
    return this.curr < this.history.length - 1;
  }

  get currentUrl(): string | null {
    if (this.curr >= 0 && this.curr < this.history.length) {
      return this.history[this.curr];
    }
    return null;
  }

  push(url: string) {
    console.log(`pushing url ${url} to history`);
    this.history = this.history.slice(0, this.curr + 1);
    this.history.push(url);
    this.curr = this.history.length - 1;
    console.log(`current index is now ${this.curr} and history length is ${this.history.length}`);
  }

  goBack() {
    if (this.canGoBack) {
      this.curr--;
    }
  }

  goForward() {
    if (this.canGoForward) {
      this.curr++;
    }
  }

  reset() {
    this.history = [];
    this.curr = 0;
  }
}

export const navigationStore = new NavigationStore();