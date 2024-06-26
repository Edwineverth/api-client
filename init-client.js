db = db.getSiblingDB('clientDB');
db.createCollection('client');
db.client.createIndex({ clientId: 1 }, { unique: true });
db.client.insertMany([
    {
        "address": "AVENIDA LAS AMERICAS CALI",
        "blocked": "SI",
        "cellPhone": "+573101234567",
        "channel": "WEB",
        "clientId": "9999999999",
        "country": "CO",
        "customerGroup": {
            "group": "01",
            "group1": "M",
            "group2": "IND",
            "group3": "FAB",
            "group4": "MUE",
            "group5": "FER"
        },
        "customerSchema": 2,
        "distrChan": "01",
        "division": "02",
        "fiscalCode1": "999999999",
        "fiscalCode2": "1234567890",
        "frequency": true,
        "frequencyDays": "VR",
        "idPortfolio": "S2",
        "immediateDelivery": true,
        "industryCode": "SUPERMERCADO",
        "industryCode1": "FARMACIA",
        "isEnrollment": true,
        "name": "SUPER TIENDA LA ECONOMICA",
        "name2": "RICARDO SUAREZ PEREZ",
        "office": "BOG90",
        "paymentCondition": "CR30",
        "paymentMethods": [
            {
                "clientId": "9999999999",
                "typeCredit": "GGCO",
                "creditLimit": 100000,
                "creditUsed": 50000,
                "amountAvailable": 50000
            },
            {
                "clientId": "9999999999",
                "typeCredit": "CO01",
                "creditLimit": 20000,
                "creditUsed": 10000,
                "amountAvailable": 10000
            },
            {
                "clientId": "9999999999",
                "typeCredit": "CO02",
                "creditLimit": 50000,
                "creditUsed": 0,
                "amountAvailable": 50000
            }
        ],
        "priceGroup": "01",
        "priceList": "02",
        "routeId": "998877",
        "salesOrg": "CO01",
        "supplyCenter": "CALI",
        "updateDate": {"date": "2024-06-04T12:05:00-0500"},
        "vendorGroup": "G99"
    }
]);
