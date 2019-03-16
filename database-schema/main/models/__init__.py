# -*- coding: utf-8 -*-
from sqlalchemy import MetaData, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import scoped_session, sessionmaker

# TODO: 環境変数で接続URI切替実装
engine = create_engine(
    "postgresql+psycopg2://artiefact@localhost/artiefact"
)
meta = MetaData(engine)
Base = declarative_base(metadata=meta)
session = scoped_session(sessionmaker(bind=engine))

# BaseにModelを登録する為にimportが必要
from .account import (
    User, Username, ProfilePicture, Profile
)  # NOQA

__all__ = [
    # account
    'User', 'Username', 'Profile', 'ProfilePicture'
]
