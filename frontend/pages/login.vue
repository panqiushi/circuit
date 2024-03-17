<script setup lang="ts">
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

const router = useRouter()
const apiFetch = $fetch.create({ baseURL: '/a' })
const gotoDashboard = () => router.push('/dashboard')

const schema = object({
    email: string().email('Invalid email').required('Required'),
    password: string()
        .min(8, 'Must be at least 8 characters')
        .required('Required')
})

type Schema = InferType<typeof schema>

const state = reactive({
    email: undefined,
    password: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
    // Do something with event.data
    console.log(event.data)
    apiFetch('/a/login', {
        method: 'POST',
        body: event.data
    }).then((res) => {
        console.log(res)
        gotoDashboard()
    })
}
</script>

<template>
    <div class="flex justify-center items-center h-screen w-full">
        <UCard class="w-1/2">
            <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
                <UFormGroup label="Email" name="email">
                    <UInput v-model="state.email" />
                </UFormGroup>

                <UFormGroup label="Password" name="password">
                    <UInput v-model="state.password" type="password" />
                </UFormGroup>

                <UButton type="submit">
                    Submit
                </UButton>
            </UForm>
        </UCard>
    </div>
</template>