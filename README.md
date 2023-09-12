# BookingToGo
test BE BookingToGo

customer
API
http://localhost:3000/customer/create Method POST
body -
{
        "nama": "dani",
        "tanggal_lahir": "1995-09-04",
        "telepon" : 85946458403,
        "kewarganegaraan": 1,
        "email": "ahmadhamdani040995@gmail.com",
        "FamilyList" : [
            {
                "nama": "atul",
                "tanggal_lahir": "1998-11-14",
                "hubungan" : "istri"
            },
            {
                "nama": "akrima",
                "tanggal_lahir": "1998-11-15",
                "hubungan" : "istri2"
            }
        ]
}

http://localhost:3000/customer/custom Method GET
reponse 
{
    "Datas": {
        "keluarga": [
            {
                "id": 9,
                "customer_id": 8,
                "hubungan": "istri",
                "nama": "atul",
                "tanggal_lahir": "1998-11-14T00:00:00Z"
            },
            {
                "id": 10,
                "customer_id": 8,
                "hubungan": "istri2",
                "nama": "akrima",
                "tanggal_lahir": "1998-11-15T00:00:00Z"
            }
        ],
        "nama": "dani",
        "tanggal_lahir": "1995-09-04T00:00:00Z",
        "telepon": 189,
        "kewarganegaraan": "Indonesia (ID)",
        "email": "ahmadhamdani040995@gmail.com"
    }
}

family_list
API
http://localhost:3000/family_list/create Method POST
body
{
    "customer_id" : 1,
    "hubungan" : "saudara",
    "nama" : "ridwan",
    "tanggal_lahir" : "1995-08-06"
}

http://localhost:3000/family_list/ Method GET
response
{
    "Datas": [
        {
            "id": 1,
            "customer_id": 4,
            "hubungan": "istri",
            "nama": "atul",
            "tanggal_lahir": "1998-11-14T00:00:00Z"
        },
        {
            "id": 2,
            "customer_id": 4,
            "hubungan": "istri2",
            "nama": "akrima",
            "tanggal_lahir": "1998-11-15T00:00:00Z"
        }
    ],
    "PaginationInfo": {
        "page": 1,
        "page_size": 10,
        "sort_by": null,
        "sort": null,
        "count": 11,
        "more_records": true
    }
}

nationality
API
http://localhost:3000/nationality/ Method GET
response 
{
    "Datas": [
        {
            "id": 1,
            "nationality_name": "Indonesia",
            "nationality_code": "(ID)"
        }
    ],
    "PaginationInfo": {
        "page": 1,
        "page_size": 10,
        "sort_by": null,
        "sort": null,
        "count": 1,
        "more_records": false
    }
}

http://localhost:3000/nationality/create Method POST
body
{
    "nationality_name" : "RUSIA" ,
    "nationality_code" : "RSA"
}