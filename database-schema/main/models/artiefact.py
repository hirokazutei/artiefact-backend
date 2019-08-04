# -*- coding: utf-8 -*-
from sqlalchemy import Date, Column, DateTime, ForeignKey
from sqlalchemy.orm import backref, relationship
from sqlalchemy.dialects.postgresql import TEXT, BIGINT, SMALLINT

from . import Base


class Artiefact(Base):
    """Artiefact"""

    __tablename__ = 'artiefact'

    id = Column(BIGINT, primary_key=True)
    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), index=True)
    created_at = Column(DateTime(timezone=True), nullable=False)
    longitude = Column(BIGINT, index=True)
    latitude = Column(BIGINT, index=True)

    user = relationship('ArtiefactUser', backref=backref('Artiefact', uselist=False))


class ArtiefactImage(Base):
    """Artiefact Image"""

    __tablename__ = 'artiefact_image'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    uri = Column(TEXT, nullable=False)
    uploaded_at = Column(DateTime(timezone=True), nullable=False)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactImage', uselist=False))


class ArtiefactProperty(Base):
    """Artiefact Property"""

    __tablename__ = 'artiefact_property'

    artiefact_id = Column(BIGINT, ForeignKey('artiefact.id'), primary_key=True)
    hint = Column(TEXT, nullable=True)
    description = Column(TEXT, nullable=True)
    artiefact_type = Column(TEXT, nullable=True)

    artiefact = relationship('Artiefact', backref=backref('ArtiefactProperty', uselist=False))


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
