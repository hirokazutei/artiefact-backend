# -*- coding: utf-8 -*-
from sqlalchemy import Column, ForeignKey
from sqlalchemy.dialects.postgresql import BIGINT, NUMERIC, ARRAY

from . import Base


class TrackingBatch(Base):
    "Batch of Tracking Data"

    __tablename__ = 'tracking_batch'

    id = Column(BIGINT, primary_key=True)
    chapter = Column(BIGINT, ForeignKey('chapter.id'), nullable=False)
    longitudes = Column("longitudes", ARRAY(NUMERIC))
    latitudes = Column("latitudes", ARRAY(NUMERIC))

