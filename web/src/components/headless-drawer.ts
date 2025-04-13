import {html, LitElement} from 'lit';
import {customElement, property} from 'lit/decorators.js';
import {getTailwindStyleSheet} from "../styles";

@customElement('headless-drawer')
export class HeadlessDrawer extends LitElement {
  // @ts-ignore
  static styles = [getTailwindStyleSheet()]

  @property()
  headerTitle = "";

  @property()
  headerSubtitle = "";

  @property({type: Boolean, reflect: true})
  open = false;

  closeDrawer() {
    this.dispatchEvent(new CustomEvent('closed'))
  }

  override render() {
    return html`
      <div class="relative z-50" aria-labelledby="slide-over-title" role="dialog" aria-modal="true">
        <!-- Background backdrop -->
        <div
            class="${this.open ? 'fixed' : ''} inset-0 bg-gray-500/75 transition-opacity ease-in-out duration-500 ${this.open ? 'opacity-100' : 'opacity-0'}"
            aria-hidden="true"></div>

        <div class="${this.open ? 'fixed' : ''} inset-0 overflow-hidden">
          <div class="absolute inset-0 overflow-hidden">
            <div class="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10 sm:pl-16">
              <div class="pointer-events-auto w-screen max-w-2xl transform transition ease-in-out duration-500 sm:duration-700 ${this.open ? 'translate-x-0' : 'translate-x-full'}">
                <div class="flex h-full flex-col divide-y divide-gray-200 bg-white shadow-xl">
                  <div class="h-0 flex-1 overflow-y-auto">
                    <div class="bg-indigo-900 px-4 py-6 sm:px-6">
                      <div class="flex items-center justify-between">
                        <h2 class="text-base font-semibold text-white" id="slide-over-title">${this.headerTitle}</h2>
                        <div class="ml-3 flex h-7 items-center">
                          <button type="button"
                                  class="relative rounded-md bg-indigo-900 text-indigo-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white" @click="${this.closeDrawer}">
                            <span class="absolute -inset-2.5"></span>
                            <span class="sr-only">Close panel</span>
                            <svg class="size-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
                                 aria-hidden="true" data-slot="icon">
                              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12"/>
                            </svg>
                          </button>
                        </div>
                      </div>
                      <div class="mt-1">
                        <p class="text-sm text-indigo-400">${this.headerSubtitle}</p>
                      </div>
                    </div>
                    <div class="flex flex-1 flex-col justify-between">
                      <slot name="body"></slot>
                    </div>
                  </div>
                  <div class="flex shrink-0 justify-end px-4 py-4">
                    <slot name="footer"></slot>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'headless-drawer': HeadlessDrawer;
  }
}
