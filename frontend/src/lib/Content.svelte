<script lang="ts">
  import type { GopherResponse, GopherItem } from "../types/gopher";
  import {getIcon, isClickableType} from "../utils/gopher";
  import ItemComponent from "./Item.svelte";

  interface Props {
    response: GopherResponse;
    onItemClick: (item: { host: string, port: string, selector: string }) => void;
  }

  let { response, onItemClick }: Props = $props();

  function handleItemClick(item: GopherItem) {
    if (isClickableType(item.type)) {
      onItemClick({
        host: item.host,
        port: item.port,
        selector: item.selector,
      });
    }
  }
</script>

<div class="gopher-content">
  {#if response.items && response.items.length > 0}
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
</style>