from sqlalchemy import MetaData, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import scoped_session, sessionmaker

engine = create_engine(
    "postgresql+psycopg2://artiefact@localhost/artiefact"
)
meta = MetaData(engine)
Base = declarative_base(metadata=meta)
session = scoped_session(sessionmaker(bind=engine))

from .account import (
    ArtiefactUser, Username, ProfilePicture, Profile
)  # NOQA
from .tokens import (
    AccessToken, AccessTokenUse
)  # NOQA
from .email import (
    RegisteredEmail, EmailVerification, EmailVerificationRequest
)
from .artiefact import (
    Artiefact,
    ArtiefactAudio,
    ArtiefactDiscovery,
    ArtiefactImage,
    ArtiefactLocation,
    ArtiefactRating,
    ArtiefactText,
    ArtiefactVideo,
)

__all__ = [
    # account
    'ArtiefactUser', 'Username', 'Profile', 'ProfilePicture',
    # token
    'AccessToken', 'AccessTokenUse',
    # email
    'RegisteredEmail', 'EmailVerification', 'EmailVerificationRequest',
    # artiefact
    'Artiefact',
    'ArtiefactAudio',
    'ArtiefactDiscovery',
    'ArtiefactImage',
    'ArtiefactLocation',
    'ArtiefactRating'
    'ArtiefactText',
    'ArtiefactVideo',
]
