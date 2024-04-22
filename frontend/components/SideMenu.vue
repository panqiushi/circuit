<script setup lang="ts">
import { Profile } from '#components';
const route = useRoute()

const links = [
  [{
    label: 'CIRCUIT',
    avatar: {
      src: ''
    },
    // badge: 200,
    config: {
      color: 'bg-slate-50',
      size: 'text-xl',
      textColor: 'text-slate-900',
      hoverColor: 'hover:bg-slate-50',
      hoverTextColor: 'hover:text-slate-900'
    }
  }, {
    label: 'Installation',
    icon: 'i-heroicons-home',
    to: '/getting-started/installation'
  }, {
    label: 'Vertical Navigation',
    icon: 'i-heroicons-chart-bar',
    to: `${route.path.startsWith('/dev') ? '/dev' : ''}/components/vertical-navigation`
  }, {
    label: 'Command Palette',
    icon: 'i-heroicons-command-line',
    to: '/components/command-palette'
  }], [{
    label: 'Examples',
    icon: 'i-heroicons-light-bulb'
  }, {
    label: 'Help',
    icon: 'i-heroicons-question-mark-circle'
  }]
]

</script>

<template>
  <Profile />
  <div class="relative h-full">
    <div class="absolute inset-y-0 left-0 h-full">
      <UVerticalNavigation :links="links" />
    </div>
  </div>



</template>