<script lang="ts">
    import Item from "./property-item.svelte";
    export let name: string
</script>

<Item name="default" common>
    <p>
        The default value of the <code>{name}</code> type. If the the default value is defined it must conform to the 
        contraints defined by the <code>{name}</code> type.
    </p>
</Item>
