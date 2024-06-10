

    
export interface EndPoint {
    name: string;

}

export interface Service {
    name: string;

    endpoints: EndPoint[]
}

export interface AppMap {
    name: string;

    services: Service[]
}
export const pagemap : AppMap = {
  "services": [
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "connection"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "transformer"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "dataset"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "column"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "mapper"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "transformation"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "processlog"
    },
    {
      "endpoints": [
        {
          "name": "read"
        },
        {
          "name": "create"
        },
        {
          "name": "update"
        },
        {
          "name": "delete"
        },
        {
          "name": "search"
        }
      ],
      "name": "importdata"
    }
  ],
  "name": "mix"
}

