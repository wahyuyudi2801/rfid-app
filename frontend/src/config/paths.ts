export const paths = {
    home: {
        getHref: () => '/'
    },
    auth: {},
    app: {
        root: {
            getHref: () => '/app'
        },
        dashboard: {
            getHref: () => '/app'
        },
        users: {
            getHref: () => '/app/users'
        },
        employees: {
            getHref: () => '/app/employees'
        },
        attendances: {
            getHref: () => '/app/attendances'
        },
        rfidCards: {
            getHref: () => '/app/rfid-cards'
        },
        locations: {
            getHref: () => '/app/locations'
        }
    }
} as const