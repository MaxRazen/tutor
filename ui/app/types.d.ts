export type Alert = {
    type: "success" | "error"
    message: string
}

export type PageData = {
    alert?: Alert
}

export type User = {
    id: string
    name: string
    email: string
    avatar: string
}

export type AuthCallbackData = {
    user: User
    authorized: boolean
    accessToken: string
}

export type AuthCallbackPageData = AuthCallbackData & PageData
