<script lang="ts">
    import Item from "./property-item.svelte";
    export let typeId: string;
</script>

<Item name="description" common>
    <p>The description of the <code>{typeId}</code> type.</p>
    <p>Inclusion rules:</p>
    <ul>
        <li>If this property's description is present it will be included.</li>
        <li>
            If this type <code>inherit</code>ed another type, the type's
            description will be included.
        </li>
    </ul>
</Item>
