<script lang="ts">
  import type { GopherItem } from "../types/gopher";
  import { isClickableType, isInfoType } from "../utils/gopher";

  interface Props {
    item: GopherItem;
    icon: string;
    onclick?: () => void;
  }

  let { item, icon, onclick }: Props = $props();

  let clickable = $derived(isClickableType(item.type));
  let info = $derived(isInfoType(item.type));

  function handleKeyPress(e: KeyboardEvent) {
    if (e.key === "Enter" && clickable && onclick) {
      onclick();
    }
  }
</script>

<div
  class="gopher-item"
  class:clickable
  class:info
  role={clickable ? "button" : "text"}
  tabindex={clickable ? 0 : undefined}
  {onclick}
  onkeypress={handleKeyPress}
>
  {#if !info && icon}
    <span class="icon">{icon}</span>
  {/if}
  <span class="display">{item.display}</span>
  {#if item.description && !info}
    <span class="description">({item.description})</span>
  {/if}
</div>

<style>
  .gopher-item {
    display: flex;
    align-items: flex-start;
    gap: 0.3rem;
    padding: 0.1rem 0;
    font-size: 1rem;
    line-height: 1.4;
  }

  .clickable {
    cursor: pointer;
    color: #0000EE;
    text-decoration: underline;
  }

  .clickable:hover {
    color: #551A8B;
  }

  .clickable:active {
    color: #FF0000;
  }

  .info {
    color: #000000;
    text-decoration: none;
  }

  .icon {
    flex-shrink: 0;
    width: 1.2rem;
    font-size: 0.9rem;
  }

  .display {
    flex: 1;
    white-space: pre-wrap;
  }

  .description {
    display: none;
  }
</style>
