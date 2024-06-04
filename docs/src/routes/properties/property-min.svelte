<script lang="ts">
    import Item from "./property-item.svelte";

    export let typeId: string;
</script>

<Item name="min">
    <p>The inclusive minimum constraint value of the {typeId}.</p>
</Item>
