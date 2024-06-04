<script lang="ts">
    import TypeCommon from "./type-common.svelte";
    import { Min, Max } from "../properties";
</script>

<TypeCommon typeId="string" builtIn>
    <svelte:fragment slot="props">
        <li>
            <Min />
        </li>
        <li>
            <Max />
        </li>
    </svelte:fragment>
</TypeCommon>
