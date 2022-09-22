<script>
	// @ts-nocheck
	import { validateEmail, validateName, validatePassword } from '@utils/validation';

	export let inpType;
	export let inpValue;
	export let isErr;
	export let handleInput;

	const LUT = {
		password: {
			label: 'Password',
			inp: {
				type: 'password',
				id: 'password-input',
				placeholder: 'Enter your password'
			},
			err: {
				condition: validatePassword,
				msg: 'Password should be minimum 8 characters'
			}
		},
		email: {
			label: 'Email',
			inp: {
				type: 'email',
				id: 'email-input',
				placeholder: 'Enter your email'
			},
			err: {
				condition: validateEmail,
				msg: 'Please enter email in correct format'
			}
		},
		name: {
			label: 'Name',
			inp: {
				type: 'text',
				id: 'name-input',
				placeholder: 'Enter your name'
			},
			err: {
				condition: validateName,
				msg: 'Add name min 2 characters'
			}
		}
	};
</script>

<label class="text-left" for={LUT[inpType].inp.id}>{LUT[inpType].label}</label>
<input
	required
	type={inpType}
	id={LUT[inpType].inp.id}
	data-label={LUT[inpType].label}
	on:change={handleInput}
	value={inpValue}
	placeholder="Enter your password"
	class={`border-black border-2 border-solid px-5 py-2 rounded ${isErr ? 'border-rose-800	' : ''}`}
/>
{#if isErr}
	<p class="text-left text-sm">{LUT[inpType].err.msg}</p>
{/if}
