import {html, LitElement} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('dialog-drawer')
export class DialogDrawer extends LitElement {
    override createRenderRoot() {
        return this;
    }

    @property({type: String})
    ref = "";

    @property({type: String})
    override title = "";

    @property({type: String})
    description = "";

    @property({type: Boolean})
    opened = false;

    handleOpen() {
        // @ts-ignore
        window.productCreationDrawer.showModal()
    }

    handleClose() {
        // @ts-ignore
        window.productCreationDrawer.close()
    }

    override render() {
        // @ts-ignore
        return html`
            <dialog id="${this.ref}"
                    class="relative z-10 opacity-0 transition-opacity duration-300 ease-in-out open:opacity-100"
                    aria-labelledby="slide-over-title" role="dialog"
                    aria-modal="true">
                <!-- Background backdrop -->
                <div class="fixed inset-0 bg-gray-500/75"></div>

                <div class="fixed inset-0 overflow-hidden">
                    <div class="absolute inset-0 overflow-hidden">
                        <div class="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10">
                            <div class="pointer-events-auto w-screen max-w-md">
                                <div class="flex h-full flex-col overflow-y-auto bg-white shadow-xl">
                                    <div class="bg-violet-900 px-4 py-6 sm:px-6">
                                        <div class="flex items-center justify-between">
                                            <h2 class="text-base font-semibold text-white" id="slide-over-title">
                                                ${this.title}
                                            </h2>
                                            <div class="ml-3 flex h-7 items-center">
                                                <button type="button"
                                                        @click="${() => this.handleClose()}"
                                                        class="relative rounded-md bg-violet-900 text-violet-200 hover:text-white focus:outline-none focus:ring-2 focus:ring-white">
                                                    <span class="absolute -inset-2.5"></span>
                                                    <span class="sr-only">Close panel</span>
                                                    <svg class="size-6" fill="none" viewBox="0 0 24 24"
                                                         stroke-width="1.5" stroke="currentColor" aria-hidden="true"
                                                         data-slot="icon">
                                                        <path stroke-linecap="round" stroke-linejoin="round"
                                                              d="M6 18 18 6M6 6l12 12"/>
                                                    </svg>
                                                </button>
                                            </div>
                                        </div>
                                        <div class="mt-1">
                                            <p class="text-sm text-violet-300">${this.description}</p>
                                        </div>
                                    </div>
                                    <div class="relative flex-1 px-4 sm:px-6">
                                        <slot name="body"></slot>
                                    </div>
                                    <div class="flex shrink-0 justify-end px-4 py-4">
                                        <slot name="footer"></slot>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </dialog>
        `;
    }
}
