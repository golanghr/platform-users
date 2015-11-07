Auth type - email, social adapter

## TODO

  - Sign In (auth type, signin properties)
  - Sign Out (user identifier)
  - Register (register type, register properties)
  - Recover (recovery type, recovery properties)
  - Info
  - Save

User
 - Id
 - Email
 - First Name
 - Last Name
 - Date Of Birth
 - Gender
 - Geo
 - Relationship (single, married, divorced)
 - Suspended
 - Deleted

 Pictures
  - Id
  - UserId
  - Picture ()
  - Default (bool)
  - Deleted

 SessionActivity
  - Id
  - UserId
  - SessionProviderId
  - Status (online, offline, idle)
  - Timestamp

 SessionProvider
  - Id
  - UserId
  - Provider (google, email, facebook, whatever)
  - Username
  - Password
  - ProviderId
  - Deleted
