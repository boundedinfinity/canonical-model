<script lang="ts">
    import Item from "./property-item.svelte";
</script>

<Item name="inherit" common>
    <p>
        Inherit from a type definition. This value of this property is a URL
        used as an ID.
    </p>

    <p>
        Some of this type's properties will override properties of the <code
            >inherit</code
        >ed types. These will be noted in the property defintions.
    </p>
</Item>
