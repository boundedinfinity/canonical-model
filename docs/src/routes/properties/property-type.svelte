<script lang="ts">
    import Item from "./property-item.svelte";
    export let name: string;
</script>

<Item name="type">
    <p>
        This <code>type</code> for this entity is <code>{name}</code>.
    </p>
    
    <p>
        If there is an <code>inherit</code> propery set, this proeprty must not be set.
    </p>
</Item>
