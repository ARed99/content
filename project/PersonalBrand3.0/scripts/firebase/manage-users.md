# Get the currently signed-in user

```
import { getAuth, onAuthStateChanged } from "firebase/auth";

const auth = getAuth();
onAuthStateChanged(auth, (user) => {
  if (user) {
    // User is signed in, see docs for a list of available properties
    // https://firebase.google.com/docs/reference/js/auth.user
    const uid = user.uid;
    // ...
  } else {
    // User is signed out
    // ...
  }
});

```

# Get a user's profile

```

import { getAuth } from "firebase/auth";

const auth = getAuth();
const user = auth.currentUser;
if (user !== null) {
  // The user object has basic properties such as display name, email, etc.
  const displayName = user.displayName;
  const email = user.email;
  const photoURL = user.photoURL;
  const emailVerified = user.emailVerified;

  // The user's ID, unique to the Firebase project. Do NOT use
  // this value to authenticate with your backend server, if
  // you have one. Use User.getToken() instead.
  const uid = user.uid;
}
```

# Get a user's provider-specific profile information

```

import { getAuth } from "firebase/auth";

const auth = getAuth();
const user = auth.currentUser;

if (user !== null) {
  user.providerData.forEach((profile) => {
    console.log("Sign-in provider: " + profile.providerId);
    console.log("  Provider-specific UID: " + profile.uid);
    console.log("  Name: " + profile.displayName);
    console.log("  Email: " + profile.email);
    console.log("  Photo URL: " + profile.photoURL);
  });
}
```

# Update a user's profile

```

import { getAuth, updateProfile } from "firebase/auth";
const auth = getAuth();
updateProfile(auth.currentUser, {
  displayName: "Jane Q. User", photoURL: "https://example.com/jane-q-user/profile.jpg"
}).then(() => {
  // Profile updated!
  // ...
}).catch((error) => {
  // An error occurred
  // ...
});

```

# Set a user's email address

```

import { getAuth, updateEmail } from "firebase/auth";
const auth = getAuth();
updateEmail(auth.currentUser, "user@example.com").then(() => {
  // Email updated!
  // ...
}).catch((error) => {
  // An error occurred
  // ...
});

```

# Send a user a verification email

```

import { getAuth, sendEmailVerification } from "firebase/auth";

const auth = getAuth();
sendEmailVerification(auth.currentUser)
  .then(() => {
    // Email verification sent!
    // ...
  });

```

# Set a user's password

```

import { getAuth, updatePassword } from "firebase/auth";

const auth = getAuth();

const user = auth.currentUser;
const newPassword = getASecureRandomPassword();

updatePassword(user, newPassword).then(() => {
  // Update successful.
}).catch((error) => {
  // An error ocurred
  // ...
});

```

# Send a password reset email

```

import { getAuth, sendPasswordResetEmail } from "firebase/auth";

const auth = getAuth();
sendPasswordResetEmail(auth, email)
  .then(() => {
    // Password reset email sent!
    // ..
  })
  .catch((error) => {
    const errorCode = error.code;
    const errorMessage = error.message;
    // ..
  });

```

# Delete a user

```

import { getAuth, deleteUser } from "firebase/auth";

const auth = getAuth();
const user = auth.currentUser;

deleteUser(user).then(() => {
  // User deleted.
}).catch((error) => {
  // An error ocurred
  // ...
});

```

# Re-authenticate a user

```

import { getAuth, reauthenticateWithCredential } from "firebase/auth";

const auth = getAuth();
const user = auth.currentUser;

// TODO(you): prompt the user to re-provide their sign-in credentials
const credential = promptForCredentials();

reauthenticateWithCredential(user, credential).then(() => {
  // User re-authenticated.
}).catch((error) => {
  // An error ocurred
  // ...
});

```
