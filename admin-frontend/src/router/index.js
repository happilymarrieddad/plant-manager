import { createRouter, createWebHistory } from 'vue-router'
import userObjs from '../../../api/pb/js/users_pb';
import userClient from '../../../api/pb/js/users_grpc_web_pb';
import uikit from 'uikit'
import store from '../store'

const client = new userClient.V1UsersClient('http://localhost:8080')

const unauthenticatedRouteNames = [
    'Login'
];

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import(/* webpackChunkName: "home" */ '../views/Home.vue')
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue')
    },
    {
        path: '/about',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    },
    {
        path: '/customers',
        name: 'Customers',
        component: () => import(/* webpackChunkName: "customerIndex" */ '../views/customers/Index.vue')
    },
    {
        path: '/customers/create',
        name: 'CustomerCreate',
        component: () => import(/* webpackChunkName: "customerCreate" */ '../views/customers/Create.vue')
    },
    {
        path: '/users',
        name: 'Users',
        component: () => import(/* webpackChunkName: "userIndex" */ '../views/users/Index.vue')
    }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
    if (unauthenticatedRouteNames.includes(to.name)) return next()

    const request = new userObjs.VerifyJWTRequest();

    console.log(window.localStorage.getItem("jwt"))
    request.setRoute(to.path);
    request.setJwt(window.localStorage.getItem("jwt"));

    client.verifyJWT(request, { 'custom-header-1': 'value1' }, (err, resp) => {
        console.log(err)
        console.log(resp)
        if (err) {
            return next({ name: 'Login' })
        }

        const isValid = resp.getValid();
        const isAuthenticated = resp.getHaspermission();
        store.commit('users/setPermissions', resp.toObject().permissionsList);

        // If JWT is NOT valid then force login
        if (!isValid) {
            uikit.notification({
                message: 'Token is invalid. Please re-login',
                status: 'danger',
                pos: 'top-center',
                timeout: 5000
            });
            window.localStorage.removeItem("jwt");
            return next({ name: 'Login' });
        }
        // If the route request is valid then continue
        if (isAuthenticated) return next();
        // Otherwise force to the only safe page
        uikit.notification({
            message: 'You are not authenticated for this route. Redirecting to Home',
            status: 'danger',
            pos: 'top-center',
            timeout: 5000
        });
        return next({ name: 'Home' });
    })

})

export default router
