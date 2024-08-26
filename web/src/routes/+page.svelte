<script>
  import { TranslationService } from "$lib/proto/api/v1beta1/translate_connect";
  import { createPromiseClient } from "@connectrpc/connect";
  import { createConnectTransport } from "@connectrpc/connect-web";
  import { onMount } from "svelte";

  onMount(async () => {
    const transport = createConnectTransport({
      baseUrl: "/",
      useBinaryFormat: true,
    });
    const client = createPromiseClient(TranslationService, transport);
    const rpcStatus = await client.healthz({});
    console.log(rpcStatus.toJson());
  });
</script>

<svelte:head>
  <title>LLM Translation</title>
</svelte:head>

<h1>LLM Translation</h1>
