<script lang="ts">
  import { onMount } from "svelte";
  import { Fetch } from "../wailsjs/go/main/App"
  import type { GopherResponse } from "./types/gopher";
  import AddressBar from "./lib/AddressBar.svelte";
  import Content from "./lib/Content.svelte";
  import {navigationStore} from "./stores/navigation.svelte";

  let url = $state("gopher://gopher.floodgap.com");
  let response = $state<GopherResponse | null>(null);
  let canGoBack = $derived(navigationStore.canGoBack);
  let canGoForward = $derived(navigationStore.canGoForward);

  onMount(() => {
    navigate(url);
  });

  async function navigate(url: string, addToHistory = true, itemType = "1") {
    const parsed = parseUrl(url);
    if (!parsed) {
      return;
    }

    response = null;

    const result = await Fetch(parsed.host, parsed.port, parsed.selector, itemType);
    if (!result.err) {
      response = result;

      if (addToHistory) {
        navigationStore.push(url);
      }
    }
  }

  function parseUrl(url: string) {
    if (url.startsWith("gopher://")) {
      url = url.substring(9);
    }

    const parts = url.split("/");
    const hostPort = parts[0].split(":");
    const host = hostPort[0];
    const port = hostPort[1] || "70";
    const selector = parts.length > 1
      ? "/" + parts.slice(1).join("/")
      : "/";

    if (!host) {
      return null;
    }

    return { host, port, selector };
  }

  function handleBack() {
    navigationStore.goBack();
    const newUrl = navigationStore.currentUrl;
    if (newUrl) {
      url = newUrl;
      navigate(newUrl, false);
    }
  }

  function handleNext() {
    navigationStore.goForward();
    const newUrl = navigationStore.currentUrl;
    if (newUrl) {
      url = newUrl;
      navigate(newUrl, false);
    }
  }

  function handleGo(newUrl: string) {
    url = newUrl;
    navigate(newUrl);
  }

  function handleItemClick(item: { host: string, port: string, selector: string, type: string }) {
    const portPart = item.port === "70" ? "" : `:${item.port}`;
    const newURL = `gopher://${item.host}${portPart}${item.selector}`;
    url = newURL;
    navigate(newURL, true, item.type);
  }
</script>

<main>
  <div class="browser">
    <header>
      <AddressBar
        bind:url={url}
        {canGoBack}
        {canGoForward}
        onGo={handleGo}
        onBack={handleBack}
        onForward={handleNext}
      />
    </header>

    <div class="content">
      {#if response}
        <Content
        {response}
         onItemClick={handleItemClick}
        />
      {/if}
    </div>
  </div>
</main>

<style>
  @font-face {
    font-family: 'Agave Nerd Font Mono';
    src: url('./assets/fonts/AgaveNerdFontMono-Regular.ttf') format('truetype');
    font-weight: normal;
    font-style: normal;
  }

  :global(body) {
    margin: 0;
    padding: 0;
    font-family: 'Agave Nerd Font Mono', monospace;
    background-color: #ffffff;
    color: #000000;
  }

  main {
    height: 100vh;
    display: flex;
    flex-direction: column;
  }

  .browser {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  header {
    background-color: #f5f5f5;
    padding: 0.75rem 1rem;
    border-bottom: 1px solid #cccccc;
  }

  h1 {
    margin: 0 0 0.75rem 0;
    font-size: 1.2rem;
    font-weight: normal;
    color: #666666;
  }

  .content {
    flex: 1;
    overflow-y: auto;
    padding: 2rem;
    background-color: #ffffff;
  }
</style>