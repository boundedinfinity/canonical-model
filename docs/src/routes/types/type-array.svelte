<script lang="ts">
    import TypeCommon from "./type-common.svelte";
</script>

<TypeCommon typeId="array" builtIn>
    <svelte:fragment slot="props">
        <li>
            <h4>items</h4>

            <p>The type information for the items in this arrays.</p>
        </li>
    </svelte:fragment>
</TypeCommon>
