import {html, LitElement} from 'lit';
import {customElement, state} from 'lit/decorators.js';
import {getTailwindStyleSheet} from "../styles";
import {encodeFormData} from "../utils";

@customElement('login-form')
export class LoginForm extends LitElement {
  // @ts-ignore
  static styles = [getTailwindStyleSheet()]

  @state()
  loginFailed = false;

  handleLogin(event: Event) {
    event.preventDefault();
    this.loginFailed = false;

    const formData = new FormData(event.target as HTMLFormElement);

    fetch('/internal/login', {
      method: 'POST',
      body: encodeFormData(formData),
      credentials: 'include',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    }).then(response => {
      if (!response.ok) {
        this.loginFailed = true;
      } else {
        window.location.href = '/homepage';
      }
    });
  }

  override render() {
    return html`
      <form class="space-y-6 group" @submit="${this.handleLogin}">
        <div>
          <label class="block text-sm font-medium leading-6 text-gray-900">
            Username
          </label>
          <div class="mt-2">
            <input
                name="username"
                type="text"
                required
                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset
                 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-900 sm:text-sm sm:leading-6 
                 ${this.loginFailed ? 'ring-red-500' : 'ring-gray-300'}"
            />
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium leading-6 text-gray-900">
            Password
          </label>
          <div class="mt-2">
            <input
                name="password"
                type="password"
                required
                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset 
                placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-900 sm:text-sm sm:leading-6 
                ${this.loginFailed ? 'ring-red-500' : 'ring-gray-300'}"
            />
          </div>
        </div>
        <span class="mt-2 text-sm text-red-500 ${this.loginFailed ? 'block' : 'hidden'}">
          Invalid username or password.
        </span>
        <button
            type="submit"
            class="flex w-full justify-center rounded-md bg-violet-900 px-3 py-1.5 text-sm font-semibold leading-6 
            text-white shadow-sm hover:bg-violet-800 focus-visible:outline focus-visible:outline-2 
            focus-visible:outline-offset-2 focus-visible:outline-violet-900
            group-invalid:pointer-events-none group-invalid:opacity-60">
          Sign In
        </button>
      </form>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'login-form': LoginForm;
  }
}
