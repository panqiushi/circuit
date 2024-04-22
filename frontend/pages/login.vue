<script setup lang="ts">
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

const router = useRouter()
const gotoDashboard = () => router.push('/dashboard')
const gotoSignup = () => router.push('/signup')

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
    console.log("event.data", event.data)
    $apiHelper('/a/login', {
        method: 'POST',
        body: event.data,
        headers: {
            'Content-Type': 'application/json'
        }
    }).then((res) => {
        console.log(res)
        gotoDashboard()
    })
}
</script>

<template>
    <div class="flex justify-center items-center h-screen w-full">
        <UCard class="w-1/3">
            <h3 class="text-2xl font-semibold text-center mb-4">Login</h3>
            <span class="text-center block mb-4">
                Don't have an account? <a @click="gotoSignup" class="text-blue-500 cursor-pointer">Signup</a></span>
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