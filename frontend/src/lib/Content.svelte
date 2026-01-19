<script lang="ts">
  import type { GopherResponse, GopherItem } from "../types/gopher";
  import {getIcon, isClickableType} from "../utils/gopher";
  import ItemComponent from "./Item.svelte";

  interface Props {
    response: GopherResponse;
    onItemClick: (item: { host: string, port: string, selector: string, type: string }) => void;
  }

  let { response, onItemClick }: Props = $props();

  function handleItemClick(item: GopherItem) {
    if (isClickableType(item.type)) {
      onItemClick({
        host: item.host,
        port: item.port,
        selector: item.selector,
        type: item.type,
      });
    }
  }
</script>

<div class="gopher-content">
  {#if response.contentType === 'image' && response.raw}
    <div class="image-container">
      <img src="data:image/gif;base64,{response.raw}" alt="Gopher image" />
    </div>
  {:else if response.contentType === 'text' && response.raw}
    <pre class="text-content">{response.raw}</pre>
  {:else if response.items && response.items.length > 0}
    <div class="items">
      {#each response.items as item, index (index)}
        <ItemComponent
          {item}
          icon={getIcon(item.type)}
          onclick={() => handleItemClick(item)}
        />
      {/each}
    </div>
  {:else}
    <div class="empty">
      <p>No items to display</p>
    </div>
  {/if}
</div>

<style>
  .gopher-content {
    max-width: 900px;
    margin: 0 auto;
    text-align: left;
  }

  .items {
    display: flex;
    flex-direction: column;
    gap: 0;
    line-height: 1.6;
  }

  .empty {
    text-align: center;
    padding: 3rem;
    color: #666;
  }

  .image-container {
    text-align: center;
    padding: 2rem;
  }

  .image-container img {
    max-width: 100%;
    height: auto;
    border: 1px solid #ccc;
  }

  .text-content {
    white-space: pre-wrap;
    font-family: inherit;
    margin: 0;
  }
</style>