<script lang="ts">
  interface Props {
    url: string;
    onGo: (url: string) => void;
  }

  let { url = $bindable(), onGo }: Props = $props();

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    onGo(url)
  }
</script>

<div class="address-bar">
  <form onsubmit={handleSubmit}>
    <input
      type="text"
      bind:value={url}
      placeholder="gopher://gopher.floodgap.com"
      aria-label="URL"
    />
    <button type="submit">Go</button>
  </form>
</div>

<style>
  .address-bar {
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }

  form {
    display: flex;
    gap: 0.5rem;
    flex: 1;
  }

  input {
    flex: 1;
    background-color: #ffffff;
    color: #000000;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 0.4rem 0.6rem;
    font-family: 'Agave Nerd Font Mono', monospace;
    font-size: 0.95rem;
  }

  input:focus {
    outline: 2px solid #0066cc;
    border-color: #0066cc;
  }

  button[type="submit"] {
    background-color: #e0e0e0;
    color: #000000;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 0.4rem 1.2rem;
    cursor: pointer;
    font-family: 'Agave Nerd Font Mono', monospace;
    font-size: 0.95rem;
    transition: background-color 0.1s;
  }

  button[type="submit"]:hover {
    background-color: #d0d0d0;
  }

  button[type="submit"]:active {
    background-color: #c0c0c0;
  }
</style>