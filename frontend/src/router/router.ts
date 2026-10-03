import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import BookingView from '../views/BookingView.vue'
import PaymentView from '../views/PaymentView.vue'
import ProductSettings from '../views/Settings/ProductSettings.vue'
import ProductSettingsSingle from '../views/Settings/ProductSettingsSingle.vue'
import ProductGroupSettings from '../views/Settings/ProductGroupSettings.vue'
import ProductGroupSettingsSingle from '../views/Settings/ProductGroupSettingsSingle.vue'
import AccountSettings from '../views/Settings/AccountSettings.vue'
import AccountSettingsSingle from '../views/Settings/AccountSettingsSingle.vue'
import CategorySettingsSingle from '../views/Settings/CategorySettingsSingle.vue'
import CategorySettings from '../views/Settings/CategorySettings.vue'
import PaymentSettings from '../views/Settings/PaymentSettings.vue'
import ClientSettings from '../views/Settings/ClientSettings.vue'
import OrderView from '../views/OrderView.vue'
import ApplicationSettings from '../views/Settings/ApplicationSettings.vue'
import UnitSettings from '../views/Settings/UnitSettings.vue'
import UnitSettingsSingle from '../views/Settings/UnitSettingsSingle.vue'
import VatSettings from '../views/Settings/VatSettings.vue'
import VatSettingsSingle from '../views/Settings/VatSettingsSingle.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/booking', component: BookingView },
  { path: '/payment', component: PaymentView },
  { name: 'ProductSettings', path: '/settings/products', component: ProductSettings},
  { name: 'ProductSettingsSingle', path: '/settings/products/:productId', component: ProductSettingsSingle},
  { name: 'ProductSettingsAdd', path: '/settings/products/add', component: ProductSettingsSingle},
  { name: 'ProductGroupSettings', path: '/settings/product-groups', component: ProductGroupSettings },
  { name: 'ProductGroupSettingsAdd', path: '/settings/product-groups/add', component: ProductGroupSettingsSingle },
  { name: 'ProductGroupSettingsEdit', path: '/settings/product-groups/:groupId', component: ProductGroupSettingsSingle },
  { name: 'CategorySettings', path: '/settings/categories/:type', component: CategorySettings },
  { name: 'CategorySettingsAdd', path: '/settings/categories/:type/add', component: CategorySettingsSingle },
  { name: 'CategorySettingsSingleEdit', path: '/settings/categories/:type/:categoryId', component: CategorySettingsSingle },
  { name: 'AccountSettings', path: '/settings/accounts', component: AccountSettings},
  { name: 'AccountSettingsSingle', path: '/settings/accounts/:accountId', component: AccountSettingsSingle},
  { name: 'AccountSettingsAdd', path: '/settings/accounts/add', component: AccountSettingsSingle},
  { name: 'UnitSettings', path: '/settings/units', component: UnitSettings },
  { name: 'UnitSettingsAdd', path: '/settings/units/add', component: UnitSettingsSingle },
  { name: 'UnitSettingsEdit', path: '/settings/units/:unitId', component: UnitSettingsSingle },
  { name: 'PaymentSettings', path: '/settings/payments', component: PaymentSettings},
  { name: 'VatSettings', path: '/settings/vats', component: VatSettings },
  { name: 'VatSettingsAdd', path: '/settings/vats/add', component: VatSettingsSingle },
  { name: 'VatSettingsEdit', path: '/settings/vats/:vatId', component: VatSettingsSingle },
  { name: 'ClientSettings', path: '/settings/client', component: ClientSettings},
  { name: 'Orders', path: '/orders/', component: OrderView},
  { path: '/settings/application', redirect: { name: 'ApplicationSettings', params: { section: 'company' } } },
  { name: 'ApplicationSettings', path: '/settings/application/:section', component: ApplicationSettings}
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router