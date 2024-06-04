<script lang="ts">
    import TypeCommon from "./type-common.svelte";
</script>

<TypeCommon typeId="object" builtIn>
    <svelte:fragment slot="props">
        <li>
            <h4>properties</h4>

            <p>The type information for the properties in this object.</p>
        </li>
    </svelte:fragment>
</TypeCommon>
