<template>
    <div class="window-height">
        <div>Contact</div>
    </div>
</template>
  
<script lang="ts">
import { reactive, toRefs, defineComponent } from 'vue';

export default defineComponent({
    name: 'ContactIndex',
    setup() {
        const state = reactive({});

        return {
            ...toRefs(state),
        };
    },
});
</script>
  