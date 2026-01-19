<script lang="ts">
  import { onMount } from "svelte";
  import { Fetch } from "../wailsjs/go/main/App"
  import type { GopherResponse } from "./types/gopher";
  import AddressBar from "./lib/AddressBar.svelte";
  import Content from "./lib/Content.svelte";
  import SearchDialog from "./lib/SearchDialog.svelte";
  import {navigationStore} from "./stores/navigation.svelte";

  let url = $state("gopher://gopher.floodgap.com");
  let response = $state<GopherResponse | null>(null);
  let canGoBack = $derived(navigationStore.canGoBack);
  let canGoForward = $derived(navigationStore.canGoForward);
  let searchItem = $state<{ host: string, port: string, selector: string, display: string } | null>(null);

  onMount(() => {
    navigate(url);
  });

  async function navigate(url: string, addToHistory = true, itemType = "1", query = "") {
    const parsed = parseUrl(url);
    if (!parsed) {
      console.log("[navigate] Failed to parse URL:", url);
      return;
    }

    response = null;

    const selector = query ? `${parsed.selector}\t${query}` : parsed.selector;
    console.log("[navigate] Fetching:", { host: parsed.host, port: parsed.port, selector, itemType, query });

    const result = await Fetch(parsed.host, parsed.port, selector, itemType);
    console.log("[navigate] Result:", result);

    if (!result.err) {
      response = result;

      if (addToHistory) {
        navigationStore.push(url, itemType);
      }
    } else {
      console.log("[navigate] Error:", result.err);
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
    const entry = navigationStore.goBack();
    if (entry) {
      url = entry.url;
      navigate(entry.url, false, entry.itemType);
    }
  }

  function handleNext() {
    const entry = navigationStore.goForward();
    if (entry) {
      url = entry.url;
      navigate(entry.url, false, entry.itemType);
    }
  }

  function handleGo(newUrl: string) {
    url = newUrl;
    navigate(newUrl);
  }

  function handleItemClick(item: { host: string, port: string, selector: string, type: string, display?: string }) {
    if (item.type === "7") {
      searchItem = {
        host: item.host,
        port: item.port,
        selector: item.selector,
        display: item.display || "Search",
      };
      return;
    }

    const portPart = item.port === "70" ? "" : `:${item.port}`;
    const newURL = `gopher://${item.host}${portPart}${item.selector}`;
    url = newURL;
    navigate(newURL, true, item.type);
  }

  function handleSearchSubmit(query: string) {
    if (!searchItem) return;

    const portPart = searchItem.port === "70" ? "" : `:${searchItem.port}`;
    const newURL = `gopher://${searchItem.host}${portPart}${searchItem.selector}`;
    url = newURL;
    searchItem = null;
    navigate(newURL, true, "1", query);
  }

  function handleSearchCancel() {
    searchItem = null;
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

  {#if searchItem}
    <SearchDialog
      prompt={searchItem.display}
      onSubmit={handleSearchSubmit}
      onCancel={handleSearchCancel}
    />
  {/if}
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