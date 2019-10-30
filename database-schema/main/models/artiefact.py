# -*- coding: utf-8 -*-
from sqlalchemy import Date, Column, DateTime, ForeignKey
from sqlalchemy.orm import backref, relationship
from sqlalchemy.dialects.postgresql import TEXT, BIGINT, SMALLINT, FLOAT

from . import Base


class Artiefact(Base):
    """Artiefact"""

    __tablename__ = 'artiefact'

    id = Column(BIGINT, primary_key=True)
    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), index=True)
    created_at = Column(DateTime(timezone=True), nullable=False)
    hint = Column(TEXT, nullable=True)
    type = Column(TEXT, nullable=True)

    user = relationship('ArtiefactUser', backref=backref('Artiefact', uselist=False))

class ArtiefactLocation(Base):
    """Artiefact Location"""

    __tablename__ = 'artiefact_location'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    longitude = Column(BIGINT, index=True)
    latitude = Column(BIGINT, index=True)


class ArtiefactText(Base):
    """Artiefact Text"""

    __tablename__ = "artiefact_text"

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    title = Column(TEXT, nullable=False)
    text = Column(TEXT, nullable=True)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactText', uselist=False))


class ArtiefactImage(Base):
    """Artiefact Image"""

    __tablename__ = 'artiefact_image'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    description = Column(TEXT, nullable=True)
    uri = Column(TEXT, nullable=False)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactImage', uselist=False))


class ArtiefactAudio(Base):
    """Artiefact Audio"""

    __tablename__ = 'artiefact_audio'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    description = Column(TEXT, nullable=True)
    duration = Column(FLOAT, nullable=False)
    uri = Column(TEXT, nullable=False)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactAudio', uselist=False))


class ArtiefactVideo(Base):
    """Artiefact Video"""

    __tablename__ = 'artiefact_video'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    description = Column(TEXT, nullable=True)
    duration = Column(FLOAT, nullable=False)
    uri = Column(TEXT, nullable=False)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactVideo', uselist=False))


class ArtiefactDiscovery(Base):
    """Artiefact Discovery"""

    __tablename__ = 'artiefact_discovery'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), index=True)
    uploaded_at = Column(DateTime(timezone=True), nullable=False)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactDiscovery', uselist=False))
    user = relationship('ArtiefactUser', backref=backref('Artiefact', uselist=False))


class ArtiefactRating(Base):
    """Artiefact Rating"""

    __tablename__ = 'artiefact_rating'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), index=True)
    rated_at = Column(DateTime(timezone=True), nullable=False)
    rating = Column(SMALLINT, nullable=False)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactDiscovery', uselist=False))
    user = relationship('ArtiefactUser', backref=backref('Artiefact', uselist=False))
