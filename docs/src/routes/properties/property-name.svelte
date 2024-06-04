<script lang="ts">
    import Item from "./property-item.svelte";
    export let typeId: string;
</script>

<Item name="name" common>
    <p>The name of the <code>{typeId}</code> type.</p>
    <p>
        Defaults to a language friendly form of the <code>type-id</code> property.
    </p>
</Item>
