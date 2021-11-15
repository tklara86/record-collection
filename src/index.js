import "./styles/scss/main.scss";


//let attention = Prompt();
const dropdown = document.querySelector('.js-nav');

dropdown.addEventListener('click', function(e) {
    const dropDownTarget = e.target.closest('.nav-item');


    if (e.target.id === 'dropDownMenuButton') {
        // Hide the dropdown
        dropDownTarget.nextElementSibling.classList.toggle('hidden');

        if (!dropDownTarget.nextElementSibling.classList.contains('hidden')) {
            e.target.style.color = '#181614';
        }

        // Toggle aria expanded attributes
        ariaToggle(e.target)

        // Click outside
        hideOnClickOutside(e.target);

        e.preventDefault();
    }

})

const hideOnClickOutside = (element) => {
    const outsideClickListener = (event) => {
        if (!element.contains(event.target)) {
            if (!!element && !!(element.offsetWidth || element.offsetHeight || element.getClientRects().length)) {
                element.nextElementSibling.classList.toggle('hidden');
                removeClickListener();
                element.style.color = '';

            }
        }
    }
    const removeClickListener = () => {
        document.removeEventListener('click', outsideClickListener);
    }
    document.addEventListener('click', outsideClickListener);
}



const ariaToggle = (toggle) => {
    toggle.getAttribute('aria-expanded') === 'false' ?
        toggle.setAttribute('aria-expanded', 'true') :
        toggle.setAttribute('aria-expanded', 'false');
}


const toggleNav = () => {
    const toggleNav = document.querySelector('.nav-brand-toggle');

    toggleNav.addEventListener('click', function (e) {
        const navContentWrapper = document.querySelector('.navContentWrapper');
        navContentWrapper.classList.toggle('hidden')
    })
}

toggleNav();


// Validate form
class FormValidator {
    constructor(form, fields) {
        this.form = form;
        this.fields = fields;
    }

    validateOnSubmit() {
        let self = this;
        this.form.addEventListener('submit', (e) => {
            e.preventDefault();
            Array.from(self.fields, field =>{
                self.validateFields(field);
            })
        })

    }

    validateOnEntry() {
        let self = this;
        Array.from(self.fields, field =>{
            const input = field.getAttribute('id')
            const inputId = document.querySelector(`#${input}`)
            inputId.addEventListener('input', e => {
                self.validateFields(field);

            })


        })
    }

    validateFields(field) {
        if (field.value.trim() === '') {
            this.setStatus(field, `${field.previousElementSibling.innerHTML} cannot be blank`, "error")
        } else {
            this.setStatus(field, null, "success")
        }
        if (field.type === 'email') {
            const re = /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
            if (re.test(field.value)) {
                this.setStatus(field, null, "success")
            } else {
                this.setStatus(field, `Please enter a valid email address`, "error")
            }
        }
        if (field.id === 'passwordConfirmation') {
            const passwordField = this.form.querySelector("#password");

            if (field.value.trim() === "") {
                this.setStatus(field, `Password confirmation required`, "error")
            } else if (field.value !== passwordField.value) {
                this.setStatus(field, `Password does not match`, "error")
            } else {
                this.setStatus(field, null, "success")
            }
        }

        if (field.id === 'endDate') {
            const startDate = this.form.querySelector('#startDate');

            const start = new Date(startDate.value)
            const end = new Date(field.value)

            if (start.getTime() >= end.getTime()) {
                this.setStatus(field, `Select date after your arrival`, "error")
            } else {
                this.setStatus(field, null, "success")
            }

        }
    }

    setStatus(field, message, status) {
        const errorMessage = field.parentElement.querySelector('.error-message');
        const formAlert = field.parentElement.querySelector('.form-alert');

        if (status === 'success') {
            formAlert.classList.remove('is-danger')
            formAlert.classList.add('is-success')
            field.classList.remove('is-danger')
            field.classList.add('is-success')
            if (errorMessage) { errorMessage.innerHTML = '' }

        }

        if (status === 'error') {
            formAlert.classList.remove('is-success')
            errorMessage.innerHTML = message
            formAlert.classList.add('is-danger')
            field.classList.add('is-danger')
        }
    }


    init() {
        this.validateOnSubmit()
        this.validateOnEntry()
    }
}

const form = document.querySelector('.js-validateForm');
const fields = document.querySelectorAll('.js-validateForm input')
const validator = new FormValidator(form, fields)

if (form != null) {
    validator.init();
}

//
// let html = `
//     <form id="checkAvailability" action="" method="post" class="js-validateForm">
//         <div class="form-inline">
//             <div class="form-control">
//                 <label class="form-label" for="startDate">Arrival</label>
//                 <input class="input-control input-control--small" type="date" name="startDate" id="startDate" placeholder="yyyy-mm-dd">
//                 <span class="form-alert"></span>
//                 <span class="error-message"></span>
//             </div>
//
//             <div class="form-control">
//                 <label class="form-label" for="endDate">Departure</label>
//                 <input class="input-control input-control--small" type="date" name="endDate" id="endDate" placeholder="yyyy-mm-dd">
//                 <span class="form-alert"></span>
//                 <span class="error-message"></span>
//             </div>
//         </div>
//     </form>
// `;
//
//
// const checkAvailability = document.querySelector('.js-checkAvailability');
// if (checkAvailability != null) {
//    checkAvailability.addEventListener('click', function () {
//         attention.custom({
//             msg: html,
//             title: 'Choose dates',
//             callback: function (result) {
//                 let form = document.getElementById('checkAvailability')
//                 let formData = new FormData(form);
//
//                 formData.append("csrf_token", "{{.CSRFToken}}")
//
//
//               fetch("/search-availability-json", {
//                     method: "post",
//                     body: formData
//                 })
//                     .then(response => response.json())
//                     .then(data => {
//                         console.log(data)
//                     })
//             }
//         }).then()
//     })
// }


function notify(msg, msgType) {
    notie.alert({
        type: msgType,
        text: msg,
    })
}

function notifyModal(title, text, icon, confirmationButton) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmationButton
    })
}


//
// // Prompt is module for alerts, notifications, and custom popup dialogs
// function Prompt() {
//     let toast = function(c) {
//        const {
//            msg = '',
//            icon =  'success',
//            position = 'top-end',
//
//        } = c
//         const Toast = Swal.mixin({
//             toast: true,
//             title: msg,
//             position: position,
//             showConfirmButton: false,
//             timer: 3000,
//             timerProgressBar: true,
//             didOpen: (toast) => {
//                 toast.addEventListener('mouseenter', Swal.stopTimer)
//                 toast.addEventListener('mouseleave', Swal.resumeTimer)
//             }
//         })
//
//         Toast.fire({
//             icon: icon
//         })
//     }
//
//     let success = function(c) {
//
//         const {
//             msg = "",
//             title = "",
//             footer =  "",
//         } = c;
//         Swal.fire({
//             icon: 'success',
//             title: title,
//             text: msg,
//             footer: footer
//         })
//     }
//
//     let error = function(c) {
//
//         const {
//             msg = "",
//             title = "",
//             footer =  "",
//         } = c;
//         Swal.fire({
//             icon: 'error',
//             title: title,
//             text: msg,
//             footer: footer
//         })
//     }
//
//    async function custom(c) {
//         const {
//             msg = "",
//             title = ""
//         } = c;
//
//         const { value: result } = await Swal.fire({
//             title: title,
//             html: msg,
//             backdrop: false,
//             focusConfirm: false,
//             showCancelButton: true,
//             preConfirm: () => {
//                 return [
//                     document.getElementById('startDate').value,
//                     document.getElementById('endDate').value
//                 ]
//             },
//             didOpen: () => {
//                 const form = document.querySelector('.js-validateForm');
//                 const fields = document.querySelectorAll('.js-validateForm input')
//                 const validator = new FormValidator(form, fields)
//
//                 if (form != null) {
//                     validator.init()
//                 }
//             }
//         })
//
//         if (result) {
//             if (result.dismiss !== Swal.DismissReason.cancel) {
//                 if (result.value !== "") {
//                     if (c.callback !== undefined) {
//                         c.callback(result);
//                     }
//                 }
//             } else {
//                 c.callback(false)
//             }
//         }
//     }
//
//     return {
//         toast: toast,
//         success: success,
//         error: error,
//         custom: custom,
//     }
// }




/* Custom select */

const customSelects = document.querySelectorAll('.custom-select');

customSelects.forEach(customSelect => {
    customSelect.addEventListener('click', (e) => {
        e.currentTarget.parentElement.classList.toggle('show');

        const option = e.currentTarget.nextElementSibling;
        const options = option.querySelectorAll('.custom-option li');
        const customOption = option.querySelector('.custom-option');

        const value = customSelect.querySelector('.value')

        options.forEach(item=> {
            item.addEventListener('click', function (e) {
                value.innerHTML = e.target.innerHTML;
                e.target.parentElement.classList.remove('show');
            })
        })

        window.addEventListener('click', function (e) {
            if(e.target !== customSelect && e.target !== customOption) {
                customSelect.parentElement.classList.remove('show');
            }
        })

    })
})
