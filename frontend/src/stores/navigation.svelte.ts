class NavigationStore {
  history = $state<{url: string, itemType: string}[]>([]);
  curr = $state(0);

  get canGoBack() {
    console.log(`cannot go back as current index is ${this.curr}`);
    return this.curr > 0;
  }

  get canGoForward() {
    console.log(`cannot go forward as current index is ${this.curr} and history length is ${this.history.length}`);
    return this.curr < this.history.length - 1;
  }

  get currentEntry() {
    if (this.curr >= 0 && this.curr < this.history.length) {
      return this.history[this.curr];
    }
    return null;
  }

  get currentUrl(): string | null {
    return this.currentEntry?.url ?? null;
  }

  push(url: string, itemType: string = "1") {
    console.log(`pushing url ${url} to history`);
    this.history = this.history.slice(0, this.curr + 1);
    this.history.push({url, itemType});
    this.curr = this.history.length - 1;
    console.log(`current index is now ${this.curr} and history length is ${this.history.length}`);
  }

  goBack() {
    if (this.canGoBack) {
      this.curr--;
    }
    return this.currentEntry;
  }

  goForward() {
    if (this.canGoForward) {
      this.curr++;
    }
    return this.currentEntry;
  }

  reset() {
    this.history = [];
    this.curr = 0;
  }
}

export const navigationStore = new NavigationStore();