                ┌────────────────────────────┐
                │         FRONTEND           │
                └────────────┬───────────────┘
                             │ HTTP Request
                             ▼
                ┌────────────────────────────┐
                │        CONTROLLER           │
                │ (ArtworkController)         │
                │ - Handles routes & params   │
                │ - Validates input           │
                │ - Returns JSON responses    │
                └────────────┬───────────────┘
                             │ Calls
                             ▼
                ┌────────────────────────────┐
                │        REPOSITORY           │
                │ (ArtworkRepository)         │
                │ - Cache lookups             │
                │ - Fetches data via registry │
                │ - Combines metadata & image │
                └────────────┬───────────────┘
                             │ Uses
                             ▼
                ┌────────────────────────────┐
                │          REGISTRY           │
                │ (IIIFRegistry)              │
                │ - Maps "source" → provider  │
                │ - Returns correct service   │
                └────────────┬───────────────┘
                             │ Provides
                             ▼
     ┌──────────────────────────────────────────────────────────┐
     │                        SERVICES                          │
     │ (ArticService, GettyService, MetService, etc.)           │
     │ - Implement IIIFProvider interface                       │
     │ - Define how to talk to each museum API                  │
     │ - Call ArtServiceClient to execute HTTP requests         │
     └────────────┬─────────────────────────────────────────────┘
                  │ Uses shared HTTP logic
                  ▼
         ┌────────────────────────────┐
         │     ArtServiceClient       │
         │  (Shared HTTP Client)      │
         │ - Manages requests, timeouts│
         │ - Builds URLs               │
         │ - Handles retries, headers  │
         └────────────────────────────┘