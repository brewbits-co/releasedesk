import {html, LitElement} from 'lit';
import {customElement} from 'lit/decorators.js';
import {getTailwindStyleSheet} from "../styles";

@customElement('error-alert')
export class ErrorAlert extends LitElement {
    // @ts-ignore
    static styles = [getTailwindStyleSheet()]

    override render() {
        return html`
          <div class="rounded-md bg-red-50 p-4">
            <div class="flex">
              <div class="shrink-0">
                <svg class="size-5 text-red-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true"
                     data-slot="icon">
                  <path fill-rule="evenodd"
                        d="M10 18a8 8 0 1 0 0-16 8 8 0 0 0 0 16ZM8.28 7.22a.75.75 0 0 0-1.06 1.06L8.94 10l-1.72 1.72a.75.75 0 1 0 1.06 1.06L10 11.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L11.06 10l1.72-1.72a.75.75 0 0 0-1.06-1.06L10 8.94 8.28 7.22Z"
                        clip-rule="evenodd"/>
                </svg>
              </div>
              <div class="ml-3">
                <h3 class="text-sm font-medium text-red-800">
                  <slot name="message"></slot>
                </h3>
                <div class="text-sm text-red-700">
                  <slot name="description"></slot>
                </div>
              </div>
            </div>
          </div>
        `;
    }
}

declare global {
    interface HTMLElementTagNameMap {
        'error-alert': ErrorAlert;
    }
}
